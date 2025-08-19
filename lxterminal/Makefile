PORTNAME=	lxterminal
PORTVERSION=	0.4.1
PORTREVISION=	2
CATEGORIES=	local sysutils

MAINTAINER=	portmaster@BSDforge.com
COMMENT=	Lightweight terminal emulator from the LXDE project
WWW=		https://github.com/lxde/lxterminal

LICENSE=	GPLv2+
LICENSE_FILE=	${WRKSRC}/COPYING

LIB_DEPENDS=	libfontconfig.so:x11-fonts/fontconfig \
		libfreetype.so:print/freetype2

USES=		cpe gmake gnome localbase pkgconfig tar:xz xorg autoreconf:2.69 # instead of pre-configure: ${SH} autogen.sh
CPE_VENDOR=	${PORTNAME}_project
USE_GNOME=	gtk30 glib20 cairo vte3 intltool
USE_XORG=	x11 xext

GNU_CONFIGURE=	yes
GNU_CONFIGURE_MANPREFIX=${PREFIX}/share
CONFIGURE_ARGS=	--enable-gtk3

PORTDOCS=	AUTHORS ChangeLog NEWS

OPTIONS_DEFINE=	DOCS NLS
OPTIONS_SUB=	yes

NLS_USES=		gettext
NLS_CONFIGURE_ENABLE=	nls
NLS_USES_OFF=		gettext-tools

USE_GITHUB=	yes
GH_ACCOUNT=	lxde

post-patch-NLS-off:
	@${REINPLACE_CMD} -e 's|src po|src|' \
		${WRKSRC}/Makefile.am

post-install-DOCS-on:
	@${MKDIR} ${STAGEDIR}${DOCSDIR}
	${INSTALL_DATA} ${PORTDOCS:S|^|${WRKSRC}/|} ${STAGEDIR}${DOCSDIR}

.include <bsd.port.mk>
