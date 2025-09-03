PORTNAME=	xwinfocus
PORTVERSION=	0.0.4
DISTVERSIONPREFIX=	v
CATEGORIES=	local

MAINTAINER=	none@freebsd.org
COMMENT=	Activate or list X11 windows, or run fallback command
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE

USES=		xorg llvm:21,build
USE_XORG=	x11 xcb xau xdmcp

BUILD_DEPENDS=	git:devel/git

USE_GITHUB=	yes
GH_ACCOUNT=	omilevskyi

# GH_PROJECT below is not expanded, so PORTNAME is used instead
GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

CFLAGS+=	-DVERSION='\"${DISTVERSION}\"' -DCOMMIT_HASH='\"${GIT_COMMIT:sh:C/(.{7}).*/\1/W}\"'

PLIST_FILES=	bin/${PORTNAME} \
		share/man/man1/${PORTNAME}.1.gz

post-install: .PHONY
	${INSTALL_MAN} ${WRKSRC}/man/${PORTNAME}.1 ${STAGEDIR}${PREFIX}/share/man/man1

.include <bsd.port.mk>
