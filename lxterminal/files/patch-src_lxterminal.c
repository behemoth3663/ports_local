--- src/lxterminal.c.orig	2025-03-31 08:33:53 UTC
+++ src/lxterminal.c
@@ -79,6 +79,7 @@ static void terminal_switch_page_event(GtkNotebook * n
 
 /* Window creation, destruction, and control. */
 static void terminal_switch_page_event(GtkNotebook * notebook, GtkWidget * page, guint num, LXTerminal * terminal);
+static gboolean terminal_button_press_event(GtkNotebook * notebook, GdkEventButton *event, LXTerminal * terminal);
 static void terminal_vte_size_allocate_event(GtkWidget *widget, GtkAllocation *allocation, Term *term);
 static void terminal_window_title_changed_event(GtkWidget * vte, Term * term);
 static gboolean terminal_close_window_confirmation_event(GtkWidget * widget, GdkEventButton * event, LXTerminal * terminal);
@@ -814,13 +815,38 @@ static void terminal_switch_page_event(GtkNotebook * n
 
         /* Propagate the title to the toplevel window. */
         const gchar * title = gtk_label_get_text(GTK_LABEL(term->label));
-        gtk_window_set_title(GTK_WINDOW(terminal->window), ((title != NULL) ? title : _("LXTerminal Terminal Emulator")));
+        gtk_window_set_title(GTK_WINDOW(terminal->window), ((title != NULL) ? title : _("LXTerminal")));
 
         /* Wait for its size to be allocated, then set its geometry */
         g_signal_connect(notebook, "size-allocate", G_CALLBACK(terminal_vte_size_allocate_event), term);
     }
 }
 
+static gboolean terminal_button_press_event(GtkNotebook * notebook, GdkEventButton * event, LXTerminal * terminal)
+{
+    if (event->button == 1) {
+        gint x, y;
+        gdk_window_get_position(event->window, &x, &y);
+        x += (gint)event->x;
+        y += (gint)event->y;
+
+        gint current_page = gtk_notebook_get_current_page(notebook);
+        GtkWidget *page = gtk_notebook_get_nth_page(notebook, current_page);
+        GtkWidget *tab_label = gtk_notebook_get_tab_label(notebook, page);
+
+        if (tab_label) {
+            GtkAllocation alloc;
+            gtk_widget_get_allocation(tab_label, &alloc);
+
+            if (alloc.x <= x && x <= alloc.x + alloc.width && alloc.y <= y && y <= alloc.y + alloc.height) {
+                return TRUE;
+            }
+        }
+    }
+
+    return FALSE;
+}
+
 /* Handler for "window-title-changed" signal on a Term. */
 static void terminal_window_title_changed_event(GtkWidget * vte, Term * term)
 {
@@ -1691,6 +1717,8 @@ LXTerminal * lxterminal_initialize(LXTermWindow * lxte
         G_CALLBACK(terminal_settings_apply), terminal);
     g_signal_connect(G_OBJECT(terminal->notebook), "switch-page", 
         G_CALLBACK(terminal_switch_page_event), terminal);
+    g_signal_connect(G_OBJECT(terminal->notebook), "button-press-event",
+        G_CALLBACK(terminal_button_press_event), terminal);
     g_signal_connect(G_OBJECT(terminal->window), "delete-event",
         G_CALLBACK(terminal_close_window_confirmation_event), terminal);
 
