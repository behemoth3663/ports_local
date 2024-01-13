PORTNAME=	gocyclo
PORTVERSION=	0.6.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local devel

MAINTAINER=	none@freebsd.org
COMMENT=	Calculate cyclomatic complexities of functions in Go source code
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	BSD3CLAUSE
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	fzipp

PLIST_FILES=	bin/${PORTNAME}

GO_TARGET=	./cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static'"

.include <bsd.port.mk>
