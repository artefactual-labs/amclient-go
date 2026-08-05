// Package amclient provides a Go client for the Archivematica Dashboard API.
//
// A Client exposes services for transfers, ingests, processing configurations,
// packages, jobs, and tasks. Each service operation returns a Response that
// wraps the underlying HTTP response.
//
// Archivematica 1.11 and newer identify the pipeline and Dashboard version in
// API response headers. The pipeline identifier is shared by Dashboard,
// MCPServer, MCPClient, and other Archivematica subsystems. Use
// Client.ServerInfo to retrieve this information explicitly, or
// Response.ServerInfo to extract it from an operation that has already
// completed. Version provides component access and comparison helpers for
// clients that need to gate behavior on an Archivematica release.
package amclient
