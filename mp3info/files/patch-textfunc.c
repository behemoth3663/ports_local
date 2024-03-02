--- textfunc.c.orig	2006-11-06 08:58:49 UTC
+++ textfunc.c
@@ -227,7 +227,7 @@ void format_output (char *format_string,mp3info *mp3, 
 
 	while((percent=strchr(format,'%'))) {
 		*percent=0;
-		printf(format);
+		printf("%s", format);
 		*percent='%';
 		code=percent+1;
 		while(*code && (code[0] != '%' && !isalpha(*code))) code++;
@@ -354,7 +354,7 @@ void format_output (char *format_string,mp3info *mp3, 
 		}
 		
 	}
-	printf(format);
+	printf("%s", format);
 }
 
 
