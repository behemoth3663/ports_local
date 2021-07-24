--- agent/plugins/configurepackage/envdetect/osdetect/linux/osdetect.go.orig	2020-11-17 21:26:34 UTC
+++ agent/plugins/configurepackage/envdetect/osdetect/linux/osdetect.go
@@ -131,12 +131,10 @@ func scanOSrelease() (string, string, error) {
 	var lines []string
 	var err error
 
-	if _, err = os.Stat("/etc/os-release"); err == nil {
-		lines, err = utils.ReadFileLines("/etc/os-release")
-	} else if _, err = os.Stat("/usr/lib/os-release"); err == nil {
-		lines, err = utils.ReadFileLines("/usr/lib/os-release")
+	if _, err = os.Stat("/var/run/os-release"); err == nil {
+		lines, err = utils.ReadFileLines("/var/run/os-release")
 	} else {
-		err = errors.New("no os-release file found")
+		err = errors.New("no /var/run/os-release file found")
 	}
 
 	if err != nil {
@@ -266,7 +264,7 @@ func parseLSBreleaseCMD(id []byte, release []byte) (st
 }
 
 func scanLSBreleaseFile() (string, string, error) {
-	lines, err := utils.ReadFileLines("/etc/lsb-release")
+	lines, err := utils.ReadFileLines("/var/run/lsb-release")
 	if err != nil {
 		return "", "", err
 	}
