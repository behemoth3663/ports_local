PORTNAME=	xwinfocus
DISTVERSION=	0.0.3
CATEGORIES=	local

MAINTAINER=	none@freebsd.org
COMMENT=	Activate (or run) X11 window
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE

USES=		xorg llvm:21,build
USE_XORG=	x11 xcb xau xdmcp

#BUILD_DEPENDS=	${LOCALBASE}/lib/libX11.a:x11/libX11 \
#		${LOCALBASE}/lib/libxcb.a:x11/libxcb \
#		${LOCALBASE}/lib/libXau.a:x11/libXau \
#		${LOCALBASE}/lib/libXdmcp.a:x11/libXdmcp

USE_GITHUB=	yes
GH_ACCOUNT=	omilevskyi
GH_TAGNAME=	beb3735

CFLAGS+=	-DVERSION='\"${DISTVERSION}\"' -DCOMMIT_HASH='\"${GH_TAGNAME}\"'

PLIST_FILES=	bin/${PORTNAME}

.include <bsd.port.mk>
