PORTNAME=	xwinfocus
DISTVERSION=	0.0.4
CATEGORIES=	local

MAINTAINER=	none@freebsd.org
COMMENT=	Activate or list X11 windows, or run fallback command
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
GH_TAGNAME=	9c2d886

CFLAGS+=	-DVERSION='\"${DISTVERSION}\"' -DCOMMIT_HASH='\"${GH_TAGNAME}\"'

PLIST_FILES=	bin/${PORTNAME} \
		share/man/man1/${PORTNAME}.1.gz

post-install: .PHONY
	${INSTALL_MAN} ${WRKSRC}/${PORTNAME}.1 ${STAGEDIR}${PREFIX}/share/man/man1

.include <bsd.port.mk>
