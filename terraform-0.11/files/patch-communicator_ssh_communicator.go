--- communicator/ssh/communicator.go.orig	2019-05-16 20:40:50 UTC
+++ communicator/ssh/communicator.go
@@ -167,6 +167,18 @@ func (c *Communicator) Connect(o terraform.UIOutput) (
 		return err
 	}
 
+	// imported from vendor/golang.org/x/crypto/ssh/common.go
+	var supportedCiphers = []string{
+		"aes128-ctr", "aes192-ctr", "aes256-ctr",
+		"aes128-gcm@openssh.com",
+		"chacha20-poly1305@openssh.com",
+		"arcfour256", "arcfour128", "arcfour",
+		"aes128-cbc",
+		"3des-cbc",
+	}
+
+	c.config.config.Ciphers = supportedCiphers
+
 	log.Printf("[DEBUG] handshaking with SSH")
 	host := fmt.Sprintf("%s:%d", c.connInfo.Host, c.connInfo.Port)
 	sshConn, sshChan, req, err := ssh.NewClientConn(c.conn, host, c.config.config)
