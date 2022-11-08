# Created by: Danilo Egea Gondolfo <danilo@FreeBSD.org>

PORTNAME=	minikube
PORTVERSION=	1.28.0
DISTVERSIONPREFIX=	v
CATEGORIES=	local sysutils

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	Run Kubernetes locally
WWW=		https://github.com/kubernetes/minikube

LICENSE=	APACHE20

BUILD_DEPENDS=	go-bindata:devel/go-bindata
RUN_DEPENDS=	kubectl:sysutils/kubectl

USES=		cpe go:modules
CPE_VENDOR=	kubernetes

USE_GITHUB=	yes
GH_ACCOUNT=	kubernetes

GO_TARGET=	./cmd/${PORTNAME}
GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
		-X k8s.io/minikube/pkg/version.version=${DISTVERSIONPREFIX}${PORTVERSION} \
		-X k8s.io/minikube/pkg/version.isoVersion=${DISTVERSIONPREFIX}${PORTVERSION} \
		-X k8s.io/minikube/pkg/version.storageProvisionerVersion=v5 \
		-X k8s.io/minikube/pkg/version.isoPath=minikube/iso"

PLIST_FILES=	bin/${PORTNAME}

SUB_FILES=	pkg-message

post-patch: .PHONY .IGNORE
	${LN} -s ${WRKSRC}/vendor/github.com/pkg/browser/browser_linux.go \
		${WRKSRC}/vendor/github.com/pkg/browser/browser_freebsd.go

pre-build: .PHONY
	cd ${WRKSRC} && \
		${LOCALBASE}/bin/go-bindata -nomemcopy \
			-o pkg/minikube/assets/assets.go \
			-pkg assets deploy/addons/... && \
		${LOCALBASE}/bin/go-bindata -nomemcopy \
			-o pkg/minikube/translate/translations.go \
			-pkg translate translations/... && \
		${LOCALBASE}/bin/gofmt${GO_SUFFIX} -s -w pkg/minikube/translate/translations.go

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
