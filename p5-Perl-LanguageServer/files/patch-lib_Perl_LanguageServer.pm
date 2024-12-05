--- lib/Perl/LanguageServer.pm.orig	2023-12-23 13:38:38 UTC
+++ lib/Perl/LanguageServer.pm
@@ -221,11 +221,16 @@ sub call_method
         $perlmethod = (defined($id)?'_rpcreq_':'_rpcnot_') . $name ;
         }
     $self -> logger ("method=$perlmethod\n") if ($debug1) ;
-    die "Unknown perlmethod $perlmethod" if (!$self -> can ($perlmethod)) ;
-
+    if ($self -> can ($perlmethod))
+        {
 no strict ;
     return $self -> $perlmethod ($workspace, $req) ;
 use strict ;
+        }
+    else
+        {
+        warn "Unknown perlmethod $perlmethod" ;
+        }
     }
 
 # ---------------------------------------------------------------------------
