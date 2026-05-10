PORTNAME=	yamlfmt
PORTVERSION=	0.21.0
DISTVERSIONPREFIX=	v
PORTREVISION=	3
CATEGORIES=	local textproc

MAINTAINER=	meta@FreeBSD.org
COMMENT=	Simple command line tool for formatting yaml files
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	google

GO_TARGET=	./cmd/${PORTNAME}

PLIST_FILES=	bin/${PORTNAME}

GIT_COMMIT=	git ls-remote \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONFULL}

CGO_ENABLED=	0
GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'main.version=${PORTVERSION}' \
	-X 'main.commit=${GIT_COMMIT:sh:C/(.{7}).*/\\1/W}'"
GO_ENV+=	GOEXPERIMENT=norandomizedheapbase64

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
