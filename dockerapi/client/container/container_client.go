// Code generated by go-swagger; DO NOT EDIT.

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new container API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for container API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ContainerArchive gets an archive of a filesystem resource in a container

Get a tar archive of a resource in the filesystem of container id.
*/
func (a *Client) ContainerArchive(params *ContainerArchiveParams) (*ContainerArchiveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerArchiveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerArchive",
		Method:             "GET",
		PathPattern:        "/containers/{id}/archive",
		ProducesMediaTypes: []string{"application/x-tar"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerArchiveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerArchiveOK), nil

}

/*
ContainerArchiveInfo gets information about files in a container

A response header `X-Docker-Container-Path-Stat` is return containing a base64 - encoded JSON object with some filesystem header information about the path.
*/
func (a *Client) ContainerArchiveInfo(params *ContainerArchiveInfoParams) (*ContainerArchiveInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerArchiveInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerArchiveInfo",
		Method:             "HEAD",
		PathPattern:        "/containers/{id}/archive",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerArchiveInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerArchiveInfoOK), nil

}

/*
ContainerAttach attaches to a container

Attach to a container to read its output or send it input. You can attach to the same container multiple times and you can reattach to containers that have been detached.

Either the `stream` or `logs` parameter must be `true` for this endpoint to do anything.

See [the documentation for the `docker attach` command](https://docs.docker.com/engine/reference/commandline/attach/) for more details.

### Hijacking

This endpoint hijacks the HTTP connection to transport `stdin`, `stdout`, and `stderr` on the same socket.

This is the response from the daemon for an attach request:

```
HTTP/1.1 200 OK
Content-Type: application/vnd.docker.raw-stream

[STREAM]
```

After the headers and two new lines, the TCP connection can now be used for raw, bidirectional communication between the client and server.

To hint potential proxies about connection hijacking, the Docker client can also optionally send connection upgrade headers.

For example, the client sends this request to upgrade the connection:

```
POST /containers/16253994b7c4/attach?stream=1&stdout=1 HTTP/1.1
Upgrade: tcp
Connection: Upgrade
```

The Docker daemon will respond with a `101 UPGRADED` response, and will similarly follow with the raw stream:

```
HTTP/1.1 101 UPGRADED
Content-Type: application/vnd.docker.raw-stream
Connection: Upgrade
Upgrade: tcp

[STREAM]
```

### Stream format

When the TTY setting is disabled in [`POST /containers/create`](#operation/ContainerCreate), the stream over the hijacked connected is multiplexed to separate out `stdout` and `stderr`. The stream consists of a series of frames, each containing a header and a payload.

The header contains the information which the stream writes (`stdout` or `stderr`). It also contains the size of the associated frame encoded in the last four bytes (`uint32`).

It is encoded on the first eight bytes like this:

```go
header := [8]byte{STREAM_TYPE, 0, 0, 0, SIZE1, SIZE2, SIZE3, SIZE4}
```

`STREAM_TYPE` can be:

- 0: `stdin` (is written on `stdout`)
- 1: `stdout`
- 2: `stderr`

`SIZE1, SIZE2, SIZE3, SIZE4` are the four bytes of the `uint32` size encoded as big endian.

Following the header is the payload, which is the specified number of bytes of `STREAM_TYPE`.

The simplest way to implement this protocol is the following:

1. Read 8 bytes.
2. Choose `stdout` or `stderr` depending on the first byte.
3. Extract the frame size from the last four bytes.
4. Read the extracted size and output it on the correct output.
5. Goto 1.

### Stream format when using a TTY

When the TTY setting is enabled in [`POST /containers/create`](#operation/ContainerCreate), the stream is not multiplexed. The data exchanged over the hijacked connection is simply the raw data from the process PTY and client's `stdin`.

*/
func (a *Client) ContainerAttach(params *ContainerAttachParams) (*ContainerAttachOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerAttachParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerAttach",
		Method:             "POST",
		PathPattern:        "/containers/{id}/attach",
		ProducesMediaTypes: []string{"application/vnd.docker.raw-stream"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerAttachReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerAttachOK), nil

}

/*
ContainerAttachWebsocket attaches to a container via a websocket
*/
func (a *Client) ContainerAttachWebsocket(params *ContainerAttachWebsocketParams) (*ContainerAttachWebsocketOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerAttachWebsocketParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerAttachWebsocket",
		Method:             "GET",
		PathPattern:        "/containers/{id}/attach/ws",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerAttachWebsocketReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerAttachWebsocketOK), nil

}

/*
ContainerChanges gets changes on a container s filesystem

Returns which files in a container's filesystem have been added, deleted,
or modified. The `Kind` of modification can be one of:

- `0`: Modified
- `1`: Added
- `2`: Deleted

*/
func (a *Client) ContainerChanges(params *ContainerChangesParams) (*ContainerChangesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerChangesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerChanges",
		Method:             "GET",
		PathPattern:        "/containers/{id}/changes",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerChangesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerChangesOK), nil

}

/*
ContainerCreate creates a container
*/
func (a *Client) ContainerCreate(params *ContainerCreateParams) (*ContainerCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerCreate",
		Method:             "POST",
		PathPattern:        "/containers/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "application/octet-stream"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerCreateCreated), nil

}

/*
ContainerDelete removes a container
*/
func (a *Client) ContainerDelete(params *ContainerDeleteParams) (*ContainerDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerDelete",
		Method:             "DELETE",
		PathPattern:        "/containers/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerDeleteNoContent), nil

}

/*
ContainerExport exports a container

Export the contents of a container as a tarball.
*/
func (a *Client) ContainerExport(params *ContainerExportParams) (*ContainerExportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerExportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerExport",
		Method:             "GET",
		PathPattern:        "/containers/{id}/export",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerExportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerExportOK), nil

}

/*
ContainerInspect inspects a container

Return low-level information about a container.
*/
func (a *Client) ContainerInspect(params *ContainerInspectParams) (*ContainerInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerInspectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerInspect",
		Method:             "GET",
		PathPattern:        "/containers/{id}/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerInspectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerInspectOK), nil

}

/*
ContainerKill kills a container

Send a POSIX signal to a container, defaulting to killing to the container.
*/
func (a *Client) ContainerKill(params *ContainerKillParams) (*ContainerKillNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerKillParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerKill",
		Method:             "POST",
		PathPattern:        "/containers/{id}/kill",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerKillReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerKillNoContent), nil

}

/*
ContainerList lists containers

Returns a list of containers. For details on the format, see [the inspect endpoint](#operation/ContainerInspect).

Note that it uses a different, smaller representation of a container than inspecting a single container. For example,
the list of linked containers is not propagated .

*/
func (a *Client) ContainerList(params *ContainerListParams) (*ContainerListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerList",
		Method:             "GET",
		PathPattern:        "/containers/json",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerListOK), nil

}

/*
ContainerLogs gets container logs

Get `stdout` and `stderr` logs from a container.

Note: This endpoint works only for containers with the `json-file` or `journald` logging driver.

*/
func (a *Client) ContainerLogs(params *ContainerLogsParams, writer io.Writer) (*ContainerLogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerLogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerLogs",
		Method:             "GET",
		PathPattern:        "/containers/{id}/logs",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerLogsReader{formats: a.formats, writer: writer},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerLogsOK), nil

}

/*
ContainerPause pauses a container

Use the cgroups freezer to suspend all processes in a container.

Traditionally, when suspending a process the `SIGSTOP` signal is used, which is observable by the process being suspended. With the cgroups freezer the process is unaware, and unable to capture, that it is being suspended, and subsequently resumed.

*/
func (a *Client) ContainerPause(params *ContainerPauseParams) (*ContainerPauseNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerPauseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerPause",
		Method:             "POST",
		PathPattern:        "/containers/{id}/pause",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerPauseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerPauseNoContent), nil

}

/*
ContainerPrune deletes stopped containers
*/
func (a *Client) ContainerPrune(params *ContainerPruneParams) (*ContainerPruneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerPruneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerPrune",
		Method:             "POST",
		PathPattern:        "/containers/prune",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerPruneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerPruneOK), nil

}

/*
ContainerRename renames a container
*/
func (a *Client) ContainerRename(params *ContainerRenameParams) (*ContainerRenameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerRenameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerRename",
		Method:             "POST",
		PathPattern:        "/containers/{id}/rename",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerRenameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerRenameNoContent), nil

}

/*
ContainerResize resizes a container t t y

Resize the TTY for a container. You must restart the container for the resize to take effect.
*/
func (a *Client) ContainerResize(params *ContainerResizeParams) (*ContainerResizeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerResizeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerResize",
		Method:             "POST",
		PathPattern:        "/containers/{id}/resize",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerResizeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerResizeOK), nil

}

/*
ContainerRestart restarts a container
*/
func (a *Client) ContainerRestart(params *ContainerRestartParams) (*ContainerRestartNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerRestartParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerRestart",
		Method:             "POST",
		PathPattern:        "/containers/{id}/restart",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerRestartReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerRestartNoContent), nil

}

/*
ContainerStart starts a container
*/
func (a *Client) ContainerStart(params *ContainerStartParams) (*ContainerStartNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerStartParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerStart",
		Method:             "POST",
		PathPattern:        "/containers/{id}/start",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerStartReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerStartNoContent), nil

}

/*
ContainerStats gets container stats based on resource usage

This endpoint returns a live stream of a container’s resource usage
statistics.

The `precpu_stats` is the CPU statistic of the *previous* read, and is
used to calculate the CPU usage percentage. It is not an exact copy
of the `cpu_stats` field.

If either `precpu_stats.online_cpus` or `cpu_stats.online_cpus` is
nil then for compatibility with older daemons the length of the
corresponding `cpu_usage.percpu_usage` array should be used.

*/
func (a *Client) ContainerStats(params *ContainerStatsParams) (*ContainerStatsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerStatsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerStats",
		Method:             "GET",
		PathPattern:        "/containers/{id}/stats",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerStatsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerStatsOK), nil

}

/*
ContainerStop stops a container
*/
func (a *Client) ContainerStop(params *ContainerStopParams) (*ContainerStopNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerStopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerStop",
		Method:             "POST",
		PathPattern:        "/containers/{id}/stop",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerStopReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerStopNoContent), nil

}

/*
ContainerTop lists processes running inside a container

On Unix systems, this is done by running the `ps` command. This endpoint is not supported on Windows.
*/
func (a *Client) ContainerTop(params *ContainerTopParams) (*ContainerTopOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerTopParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerTop",
		Method:             "GET",
		PathPattern:        "/containers/{id}/top",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerTopReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerTopOK), nil

}

/*
ContainerUnpause unpauses a container

Resume a container which has been paused.
*/
func (a *Client) ContainerUnpause(params *ContainerUnpauseParams) (*ContainerUnpauseNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerUnpauseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerUnpause",
		Method:             "POST",
		PathPattern:        "/containers/{id}/unpause",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerUnpauseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerUnpauseNoContent), nil

}

/*
ContainerUpdate updates a container

Change various configuration options of a container without having to recreate it.
*/
func (a *Client) ContainerUpdate(params *ContainerUpdateParams) (*ContainerUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerUpdate",
		Method:             "POST",
		PathPattern:        "/containers/{id}/update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerUpdateOK), nil

}

/*
ContainerWait waits for a container

Block until a container stops, then returns the exit code.
*/
func (a *Client) ContainerWait(params *ContainerWaitParams) (*ContainerWaitOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContainerWaitParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ContainerWait",
		Method:             "POST",
		PathPattern:        "/containers/{id}/wait",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ContainerWaitReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ContainerWaitOK), nil

}

/*
PutContainerArchive extracts an archive of files or folders to a directory in a container

Upload a tar archive to be extracted to a path in the filesystem of container id.
*/
func (a *Client) PutContainerArchive(params *PutContainerArchiveParams) (*PutContainerArchiveOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutContainerArchiveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutContainerArchive",
		Method:             "PUT",
		PathPattern:        "/containers/{id}/archive",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/octet-stream", "application/x-tar"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutContainerArchiveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutContainerArchiveOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
