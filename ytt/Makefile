PORTNAME=	ytt
PORTVERSION=	0.53.2
DISTVERSIONPREFIX=	v
CATEGORIES=	local textproc
PKGNAMESUFFIX=	-local

MAINTAINER=	none@freebsd.org
COMMENT=	YAML tool with templates and user provided values
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	k14s

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'github.com/vmware-tanzu/carvel-ytt/pkg/version.Version=${PORTVERSION}'"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
