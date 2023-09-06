PORTNAME=	eksctl
PORTVERSION=	0.155.0
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	CLI for Amazon EKS
WWW=		https://eksctl.io/

LICENSE=	APACHE20

#RUN_DEPENDS=	kubectl:sysutils/kubectl

USES=		go:modules

PORTSCOUT=	limit:^[0-9\.]*$$

USE_GITHUB=	yes
GH_ACCOUNT=	weaveworks

PLIST_FILES=	bin/${PORTNAME}

GO_TARGET=	./cmd/${PORTNAME}

#GO_BUILDFLAGS=	-ldflags="-X github.com/weaveworks/eksctl/pkg/version.gitCommit=97ec4b2d946c
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
		-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/pkg/version.gitCommit=${PORTVERSION} \
		-X github.com/${GH_ACCOUNT}/${GH_PROJECT}/pkg/version.builtDate=${%FT%TZ:L:gmtime}"

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"
.include <bsd.port.mk>
