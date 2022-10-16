# Created by: Danilo Egea Gondolfo <danilo@FreeBSD.org>

PORTNAME=	kops
PORTVERSION=	1.25.1
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	Kubernetes Operations

LICENSE=	APACHE20

BUILD_DEPENDS=	go-bindata:devel/go-bindata

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	kubernetes

GO_PKGNAME=	k8s.io/${PORTNAME}
GO_TARGET=	${GO_PKGNAME}/cmd/${PORTNAME}
GO_BUILDFLAGS+=	-ldflags="-s -w -X k8s.io/kops.Version=${PORTVERSION}"

PLIST_FILES=	bin/${PORTNAME}

pre-build: .PHONY
	cd ${WRKSRC} && \
		${LOCALBASE}/bin/go-bindata \
			-o upup/models/bindata.go -pkg models \
			-prefix upup/models/ upup/models/...

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
