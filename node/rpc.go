package node

// configureRPC is a helper method to configure all the various RPC endpoints during node
// startup. It's not meant to be called at any time afterwards as it makes certain
// assumptions about the state of the node.
func (node *Node) startRPC() error {

	if node.config.RPC.EnableHTTPS {
		// Configure Secure HTTPS.
		if node.config.RPC.TLS.Certificate != "" && node.config.RPC.TLS.Key != "" {
			if err := node.https.setListenAddr(node.config.RPC.HTTPSHost, node.config.RPC.HTTPSPort); err != nil {
				return err
			}
			config := httpConfig{
				CorsAllowedOrigins: node.config.RPC.HTTPCors,
				Vhosts:             node.config.RPC.HTTPSVirtualHosts,
				Modules:            node.config.RPC.Endpoints,
				prefix:             "",
			}
			if err := node.https.enableRPC(node.rpcAPIs, config); err != nil {
				return err
			}
			if err := node.https.startSecure(node.config.RPC.TLS.Certificate, node.config.RPC.TLS.Key); err != nil {
				return err
			}
		} else {
			return ErrNodeStopped
		}
	}

	if node.config.RPC.EnableHTTP {
		// Configure HTTP.
		if node.config.RPC.HTTPHost != "" {
			config := httpConfig{
				CorsAllowedOrigins: node.config.RPC.HTTPCors,
				Vhosts:             node.config.RPC.HTTPVirtualHosts,
				Modules:            node.config.RPC.Endpoints,
				prefix:             "",
			}
			if err := node.http.setListenAddr(node.config.RPC.HTTPHost, node.config.RPC.HTTPPort); err != nil {
				return err
			}
			if err := node.http.enableRPC(node.rpcAPIs, config); err != nil {
				return err
			}
		}
		if err := node.http.start(); err != nil {
			return err
		}
	}

	if node.config.RPC.EnableWSS {
		// Configure Secure WebSocket.
		if node.config.RPC.TLS.Certificate != "" && node.config.RPC.TLS.Key != "" {
			if node.config.RPC.WSSHost != "" {
				server := node.wssServerForPort(node.config.RPC.WSSPort)
				config := wsConfig{
					Modules: node.config.RPC.Endpoints,
					Origins: node.config.RPC.WSOrigins,
					prefix:  "",
				}
				if err := server.setListenAddr(node.config.RPC.WSSHost, node.config.RPC.WSSPort); err != nil {
					return err
				}
				if err := server.enableWS(node.rpcAPIs, config); err != nil {
					return err
				}
			}
			if err := node.wss.startSecure(node.config.RPC.TLS.Certificate, node.config.RPC.TLS.Key); err != nil {
				return err
			}
		} else {
			return ErrNodeStopped
		}
	}

	if node.config.RPC.EnableWS {
		// Configure WebSocket.
		if node.config.RPC.WSHost != "" {
			server := node.wsServerForPort(node.config.RPC.WSPort)
			config := wsConfig{
				Modules: node.config.RPC.Endpoints,
				Origins: node.config.RPC.WSOrigins,
				prefix:  "",
			}
			if err := server.setListenAddr(node.config.RPC.WSHost, node.config.RPC.WSPort); err != nil {
				return err
			}
			if err := server.enableWS(node.rpcAPIs, config); err != nil {
				return err
			}
		}
		if err := node.ws.start(); err != nil {
			return err
		}
	}

	return nil
}

func (node *Node) wsServerForPort(port int) *httpServer {
	if node.config.RPC.HTTPHost == "" || node.http.port == port {
		return node.http
	}
	return node.ws
}

func (node *Node) wssServerForPort(port int) *httpServer {
	if node.config.RPC.HTTPSHost == "" || node.https.port == port {
		return node.https
	}
	return node.wss
}

func (node *Node) stopRPC() {
	node.http.stop()
	node.https.stop()
	node.ws.stop()
	node.wss.stop()
}
