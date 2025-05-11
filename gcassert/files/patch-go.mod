--- go.mod.orig	2025-04-30 16:46:44 UTC
+++ go.mod
@@ -1,8 +1,15 @@ module github.com/jordanlewis/gcassert
 module github.com/jordanlewis/gcassert
 
-go 1.13
+go 1.24
 
 require (
 	github.com/stretchr/testify v1.6.1
 	golang.org/x/tools v0.17.0
+)
+
+require (
+	github.com/davecgh/go-spew v1.1.0 // indirect
+	github.com/pmezard/go-difflib v1.0.0 // indirect
+	golang.org/x/mod v0.14.0 // indirect
+	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
 )
