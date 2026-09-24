--- internal/command/state_migrate.go.orig	2026-09-23 11:57:56 UTC
+++ internal/command/state_migrate.go
@@ -465,9 +465,6 @@ func (c *StateMigrateCommand) getSingleProvider(ctx co
 // - Potential for downloading different versions of the same provider
 // - Need to keep the locks separate for source and destination providers; destination providers are added to the dependency lock file.
 func (c *StateMigrateCommand) getSingleProvider(ctx context.Context, storeName string, reqs providerreqs.Requirements, locks *depsfile.Locks, upgrade bool, location string, view views.StateMigrate) (output bool, resultingLock *depsfile.Locks, diags tfdiags.Diagnostics) {
-	ctx, span := tracer.Start(ctx, "install state migration "+location+" provider")
-	defer span.End()
-
 	// We expect to download only one provider
 	if len(reqs) != 1 {
 		panic(fmt.Sprintf("expected exactly one provider requirement for the destination state store provider, got %d", len(reqs)))
