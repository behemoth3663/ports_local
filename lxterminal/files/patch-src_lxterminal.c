--- src/lxterminal.c.orig	2025-03-31 08:33:53 UTC
+++ src/lxterminal.c
@@ -79,6 +79,7 @@ static void terminal_switch_page_event(GtkNotebook * n
 
 /* Window creation, destruction, and control. */
 static void terminal_switch_page_event(GtkNotebook * notebook, GtkWidget * page, guint num, LXTerminal * terminal);
+static gboolean terminal_button_press_event(GtkNotebook * notebook, GdkEventButton *event, LXTerminal * terminal);
 static void terminal_vte_size_allocate_event(GtkWidget *widget, GtkAllocation *allocation, Term *term);
 static void terminal_window_title_changed_event(GtkWidget * vte, Term * term);
 static gboolean terminal_close_window_confirmation_event(GtkWidget * widget, GdkEventButton * event, LXTerminal * terminal);
@@ -290,24 +291,30 @@ static void terminal_initialize_switch_tab_accelerator
  * These switch to the tab selected by the digit, if it exists. */
 static void terminal_initialize_switch_tab_accelerator(Term * term)
 {
-    if ((term->index + 1) < 10)
-    {
-        /* Formulate the accelerator name. */
-        char switch_tab_accel[1 + 3 + 1 + 1 + 1]; /* "<ALT>n" */
-        sprintf(switch_tab_accel, "<ALT>%d", term->index + 1);
+    guint key;
+    GdkModifierType mods = GDK_MOD1_MASK;
+    term->closure = g_cclosure_new_swap(G_CALLBACK(terminal_switch_tab_accelerator), term, NULL);
 
-        /* Parse the accelerator name. */
-        guint key;
-        GdkModifierType mods;
-        gtk_accelerator_parse(switch_tab_accel, &key, &mods);
-
-        /* Define the accelerator. */
-        term->closure = g_cclosure_new_swap(G_CALLBACK(terminal_switch_tab_accelerator), term, NULL);
-        if (gtk_accel_group_from_accel_closure(term->closure) == NULL)
-        {
-            gtk_accel_group_connect(term->parent->accel_group, key, mods, GTK_ACCEL_LOCKED, term->closure);
-        }
+    if (term->index < 10)
+    {
+        key = GDK_KEY_0 + (term->index + 1) % 10;
     }
+    else if (term->index < 20)
+    {
+        mods |= GDK_SHIFT_MASK;
+        static const guint shifted_keys[] = {
+            GDK_KEY_exclam, GDK_KEY_at, GDK_KEY_numbersign, GDK_KEY_dollar,
+            GDK_KEY_percent, GDK_KEY_asciicircum, GDK_KEY_ampersand, GDK_KEY_asterisk,
+            GDK_KEY_parenleft, GDK_KEY_parenright
+        };
+        key = shifted_keys[term->index % 10];
+    }
+    else
+    {
+        return;
+    }
+
+    gtk_accel_group_connect(term->parent->accel_group, key, mods, GTK_ACCEL_LOCKED, term->closure);
 }
 
 /* update <ALT>n status. */
@@ -814,13 +821,38 @@ static void terminal_switch_page_event(GtkNotebook * n
 
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
@@ -1691,6 +1723,8 @@ LXTerminal * lxterminal_initialize(LXTermWindow * lxte
         G_CALLBACK(terminal_settings_apply), terminal);
     g_signal_connect(G_OBJECT(terminal->notebook), "switch-page", 
         G_CALLBACK(terminal_switch_page_event), terminal);
+    g_signal_connect(G_OBJECT(terminal->notebook), "button-press-event",
+        G_CALLBACK(terminal_button_press_event), terminal);
     g_signal_connect(G_OBJECT(terminal->window), "delete-event",
         G_CALLBACK(terminal_close_window_confirmation_event), terminal);
 
