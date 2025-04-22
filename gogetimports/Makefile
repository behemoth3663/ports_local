PORTNAME=	gogetimports
PORTVERSION=	0.0.1
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Get JSON-formatted list of Go imports
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE

#USES=		go:modules
USES=		go

USE_GITHUB=	yes
GH_ACCOUNT=	jgautheron
GH_TAGNAME=	dd38777

PLIST_FILES=	bin/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags -static"

do-build: .PHONY
	(cd ${GO_WRKSRC} && \
		${SETENV} ${MAKE_ENV} ${GO_ENV} ${GO_CMD} build ${GO_BUILDFLAGS} \
			-o ${GO_WRKDIR_BIN}/${PORTNAME})

.include <bsd.port.mk>
