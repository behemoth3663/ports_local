PORTNAME=	terrascan
PORTVERSION=	1.19.9
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local security

MAINTAINER=	none@freebsd.org
COMMENT=	Detect compliance and security violations across IaC
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT} https://runterrascan.io/

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
USE_GITLAB=	nodefault

GH_ACCOUNT=	tenable

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

CGO_ENABLED=	0
# https://protobuf.dev/reference/go/faq/#namespace-conflict
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X google.golang.org/protobuf/reflect/protoregistry.conflictPolicy=ignore"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
