--- internal/communicator/ssh/communicator.go.orig	2019-12-02 18:18:41 UTC
+++ internal/communicator/ssh/communicator.go
@@ -184,6 +184,16 @@ func (c *Communicator) Connect(o terraform.UIOutput) (
 		return err
 	}
 
+	// imported from vendor/golang.org/x/crypto/ssh/common.go
+	c.config.config.Ciphers = []string{
+		"aes128-ctr", "aes192-ctr", "aes256-ctr",
+		"aes128-gcm@openssh.com", "aes256-gcm@openssh.com",
+		"chacha20-poly1305@openssh.com",
+		"arcfour256", "arcfour128", "arcfour",
+		"aes128-cbc",
+		"3des-cbc",
+	}
+
 	log.Printf("[DEBUG] Connection established. Handshaking for user %v", c.connInfo.User)
 	sshConn, sshChan, req, err := ssh.NewClientConn(c.conn, hostAndPort, c.config.config)
 	if err != nil {
