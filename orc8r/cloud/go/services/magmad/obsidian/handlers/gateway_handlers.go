/*
 * Copyright (c) Facebook, Inc. and its affiliates.
 * All rights reserved.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 */

package handlers

import (
	"io"
	"net/http"

	"magma/orc8r/cloud/go/datastore"
	models2 "magma/orc8r/cloud/go/models"
	"magma/orc8r/cloud/go/obsidian"
	"magma/orc8r/cloud/go/pluginimpl/handlers"
	"magma/orc8r/cloud/go/services/magmad"
	magmad_models "magma/orc8r/cloud/go/services/magmad/obsidian/models"
	"magma/orc8r/lib/go/protos"

	"github.com/labstack/echo"
)

const (
	CommandRootV1           = handlers.ManageGatewayPath + "/command"
	RebootGatewayV1         = CommandRootV1 + "/reboot"
	RestartServicesV1       = CommandRootV1 + "/restart_services"
	GatewayPingV1           = CommandRootV1 + "/ping"
	GatewayGenericCommandV1 = CommandRootV1 + "/generic"
	TailGatewayLogsV1       = CommandRootV1 + "/tail_logs"
)

func rebootGateway(c echo.Context) error {
	networkID, gatewayID, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}

	err := magmad.GatewayReboot(networkID, gatewayID)
	if err != nil {
		if datastore.IsErrNotFound(err) {
			return obsidian.HttpError(err, http.StatusNotFound)
		}
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func restartServices(c echo.Context) error {
	networkID, gatewayID, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}

	var services []string
	err := c.Bind(&services)
	if err != nil {
		return obsidian.HttpError(err, http.StatusBadRequest)
	}
	err = magmad.GatewayRestartServices(networkID, gatewayID, services)
	if err != nil {
		if datastore.IsErrNotFound(err) {
			return obsidian.HttpError(err, http.StatusNotFound)
		}
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}

	return c.NoContent(http.StatusOK)
}

func gatewayPing(c echo.Context) error {
	networkID, gatewayID, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}

	pingRequest := magmad_models.PingRequest{}
	err := c.Bind(&pingRequest)
	response, err := magmad.GatewayPing(networkID, gatewayID, pingRequest.Packets, pingRequest.Hosts)
	if err != nil {
		if datastore.IsErrNotFound(err) {
			return obsidian.HttpError(err, http.StatusNotFound)
		}
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}
	var pingResponse magmad_models.PingResponse
	for _, ping := range response.Pings {
		pingResult := &magmad_models.PingResult{
			HostOrIP:           &ping.HostOrIp,
			NumPackets:         &ping.NumPackets,
			PacketsTransmitted: ping.PacketsTransmitted,
			PacketsReceived:    ping.PacketsReceived,
			AvgResponseMs:      ping.AvgResponseMs,
			Error:              ping.Error,
		}
		pingResponse.Pings = append(pingResponse.Pings, pingResult)
	}
	return c.JSON(http.StatusOK, &pingResponse)
}

func gatewayGenericCommand(c echo.Context) error {
	networkID, gatewayID, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}

	request := magmad_models.GenericCommandParams{}
	err := c.Bind(&request)
	if err != nil {
		return obsidian.HttpError(err, http.StatusBadRequest)
	}
	params, err := models2.JSONMapToProtobufStruct(request.Params)
	if err != nil {
		return obsidian.HttpError(err, http.StatusBadRequest)
	}
	genericCommandParams := protos.GenericCommandParams{
		Command: *request.Command,
		Params:  params,
	}

	response, err := magmad.GatewayGenericCommand(networkID, gatewayID, &genericCommandParams)
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}

	resp, err := models2.ProtobufStructToJSONMap(response.Response)
	genericCommandResponse := magmad_models.GenericCommandResponse{
		Response: resp,
	}
	return c.JSON(http.StatusOK, genericCommandResponse)
}

func tailGatewayLogs(c echo.Context) error {
	networkID, gatewayID, nerr := obsidian.GetNetworkAndGatewayIDs(c)
	if nerr != nil {
		return nerr
	}

	request := magmad_models.TailLogsRequest{}
	err := c.Bind(&request)
	if err != nil {
		return obsidian.HttpError(err, http.StatusBadRequest)
	}

	stream, err := magmad.TailGatewayLogs(networkID, gatewayID, request.Service)
	if err != nil {
		return obsidian.HttpError(err, http.StatusInternalServerError)
	}

	go func() {
		<-c.Request().Context().Done()
	}()
	// https://echo.labstack.com/cookbook/streaming-response
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlainCharsetUTF8)
	c.Response().Header().Set(echo.HeaderXContentTypeOptions, "nosniff")
	c.Response().WriteHeader(http.StatusOK)
	for {
		line, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return obsidian.HttpError(err, http.StatusInternalServerError)
		}

		if _, err := c.Response().Write([]byte(line.Line)); err != nil {
			return obsidian.HttpError(err, http.StatusInternalServerError)
		}
		c.Response().Flush()
	}

	return c.NoContent(http.StatusNoContent)
}
