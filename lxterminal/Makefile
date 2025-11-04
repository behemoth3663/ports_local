PORTNAME=	lxterminal
PORTVERSION=	0.4.1
PORTREVISION=	4
CATEGORIES=	local sysutils

MAINTAINER=	portmaster@BSDforge.com
COMMENT=	Lightweight terminal emulator from the LXDE project
WWW=		https://github.com/lxde/lxterminal

LICENSE=	GPLv2+
LICENSE_FILE=	${WRKSRC}/COPYING

LIB_DEPENDS=	libfontconfig.so:x11-fonts/fontconfig \
		libfreetype.so:print/freetype2

USES=		cpe gmake gnome localbase pkgconfig tar:xz xorg autoreconf:2.69 llvm:21,build # instead of pre-configure: ${SH} autogen.sh
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

CFLAGS:=        -O3 ${CFLAGS:N-O*}
CFLAGS+=	-fomit-frame-pointer -fno-unwind-tables -fno-asynchronous-unwind-tables

CFLAGS+=	-fwhole-program-vtables
CFLAGS+=	-fno-unique-section-names

CFLAGS+=	-flto=full
LDFLAGS+=	-Wl,--lto-O3

CFLAGS+=	-faddrsig
LDFLAGS+=	-Wl,--icf=all
#LDFLAGS+=	-Wl,--print-icf-sections

CFLAGS+=	-ffunction-sections
CFLAGS+=	-fdata-sections
LDFLAGS+=	-Wl,--gc-sections

#LDFLAGS+=	-Wl,--hash-style=gnu
#LDFLAGS+=	-fuse-ld=lld

#no compile LDFLAGS+=	-Wl,--pack-dyn-relocs=relr

#no effect
LDFLAGS+=	-Wl,--as-needed
LDFLAGS+=	-Wl,--build-id=none
LDFLAGS+=	-Wl,--sort-common

post-patch-NLS-off:
	@${REINPLACE_CMD} -e 's|src po|src|' \
		${WRKSRC}/Makefile.am

post-install-DOCS-on:
	@${MKDIR} ${STAGEDIR}${DOCSDIR}
	${INSTALL_DATA} ${PORTDOCS:S|^|${WRKSRC}/|} ${STAGEDIR}${DOCSDIR}

.include <bsd.port.mk>
