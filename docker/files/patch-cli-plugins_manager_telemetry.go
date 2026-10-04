--- cli-plugins/manager/telemetry.go.orig	2026-09-23 09:45:10 UTC
+++ cli-plugins/manager/telemetry.go
@@ -7,9 +7,6 @@ import (
 
 	"github.com/docker/cli/cli-plugins/metadata"
 	"github.com/spf13/cobra"
-	"go.opentelemetry.io/otel"
-	"go.opentelemetry.io/otel/attribute"
-	"go.opentelemetry.io/otel/baggage"
 )
 
 const (
@@ -23,46 +20,19 @@ const (
 	//
 	// It is a copy of the const defined in [command.dockerCLIAttributePrefix].
 	dockerCLIAttributePrefix = "docker.cli."
-	cobraCommandPath         = attribute.Key("cobra.command_path")
+	cobraCommandPath         = "cobra.command_path"
 )
 
-func getPluginResourceAttributes(cmd *cobra.Command, plugin Plugin) attribute.Set {
+func getPluginResourceAttributes(cmd *cobra.Command, plugin Plugin) []string {
 	commandPath := cmd.Annotations[metadata.CommandAnnotationPluginCommandPath]
 	if commandPath == "" {
 		commandPath = fmt.Sprintf("%s %s", cmd.CommandPath(), plugin.Name)
 	}
-
-	attrSet := attribute.NewSet(
-		cobraCommandPath.String(commandPath),
-	)
-
-	kvs := make([]attribute.KeyValue, 0, attrSet.Len())
-	for iter := attrSet.Iter(); iter.Next(); {
-		attr := iter.Attribute()
-		kvs = append(kvs, attribute.KeyValue{
-			Key:   dockerCLIAttributePrefix + attr.Key,
-			Value: attr.Value,
-		})
-	}
-	return attribute.NewSet(kvs...)
+	return []string{dockerCLIAttributePrefix + cobraCommandPath + "=" + commandPath}
 }
 
 func appendPluginResourceAttributesEnvvar(env []string, cmd *cobra.Command, plugin Plugin) []string {
-	if attrs := getPluginResourceAttributes(cmd, plugin); attrs.Len() > 0 {
-		// Construct baggage members for each of the attributes.
-		// Ignore any failures as these aren't significant and
-		// represent an internal issue.
-		members := make([]baggage.Member, 0, attrs.Len())
-		for iter := attrs.Iter(); iter.Next(); {
-			attr := iter.Attribute()
-			m, err := baggage.NewMemberRaw(string(attr.Key), attr.Value.AsString())
-			if err != nil {
-				otel.Handle(err)
-				continue
-			}
-			members = append(members, m)
-		}
-
+	if attrs := getPluginResourceAttributes(cmd, plugin); len(attrs) > 0 {
 		// Combine plugin added resource attributes with ones found in the environment
 		// variable. Our own attributes should be namespaced so there shouldn't be a
 		// conflict. We do not parse the environment variable because we do not want
@@ -71,11 +41,7 @@ func appendPluginResourceAttributesEnvvar(env []string
 		if v := strings.TrimSpace(os.Getenv(resourceAttributesEnvVar)); v != "" {
 			attrsSlice = append(attrsSlice, v)
 		}
-		if b, err := baggage.New(members...); err != nil {
-			otel.Handle(err)
-		} else if b.Len() > 0 {
-			attrsSlice = append(attrsSlice, b.String())
-		}
+		attrsSlice = append(attrsSlice, strings.Join(attrs, ","))
 
 		if len(attrsSlice) > 0 {
 			env = append(env, resourceAttributesEnvVar+"="+strings.Join(attrsSlice, ","))
