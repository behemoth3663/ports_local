PORTNAME=	devhelp
PORTVERSION=	43.0
CATEGORIES=	devel gnome
MASTER_SITES=	GNOME
PKGNAMESUFFIX=	-local
DIST_SUBDIR=	gnome

MAINTAINER=	gnome@FreeBSD.org
COMMENT=	Developer help browser for GNOME
WWW=		https://wiki.gnome.org/Apps/Devhelp

LICENSE=	GPLv3
LICENSE_FILE=	${WRKSRC}/LICENSES/GPL-3.0-or-later.txt

BUILD_DEPENDS=	itstool:textproc/itstool
LIB_DEPENDS=	libsoup-2.4.so:devel/libsoup \
		libwebkit2gtk-4.1.so:www/webkit2-gtk@41

USES=		compiler:c++11-lib gettext gnome localbase meson pkgconfig \
		python:build tar:xz
USE_GNOME=	cairo gtk30 introspection

PORTSCOUT=	limitw:1,even

USE_LDCONFIG=	yes

GLIB_SCHEMAS=	org.gnome.devhelp.gschema.xml \
		org.gnome.libdevhelp-3.gschema.xml

.include <bsd.port.mk>
