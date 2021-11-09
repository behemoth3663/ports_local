# Created by: Danilo Egea Gondolfo <danilo@FreeBSD.org>

PORTNAME=	minikube
PORTVERSION=	1.24.0
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	Run Kubernetes locally

LICENSE=	APACHE20

BUILD_DEPENDS=	go-bindata:devel/go-bindata
RUN_DEPENDS=	kubectl:sysutils/kubectl

USES=		cpe go:modules
CPE_VENDOR=	kubernetes

USE_GITHUB=	yes
GH_ACCOUNT=	kubernetes

GO_TARGET=	./cmd/${PORTNAME}
GO_BUILDFLAGS=	-ldflags="-X k8s.io/minikube/pkg/version.version=${DISTVERSIONPREFIX}${PORTVERSION} \
		-X k8s.io/minikube/pkg/version.isoVersion=${DISTVERSIONPREFIX}${PORTVERSION:R:R}.${PORTVERSION:R:E}.0 \
		-X k8s.io/minikube/pkg/version.storageProvisionerVersion=v4 \
		-X k8s.io/minikube/pkg/version.isoPath=minikube/iso"

PLIST_FILES=	bin/${PORTNAME}

SUB_FILES=	pkg-message

post-patch:
	cd ${WRKSRC}/vendor/github.com/pkg/browser && \
		${LN} -s browser_linux.go browser_freebsd.go || :

pre-build:
	cd ${WRKSRC} && \
		${LOCALBASE}/bin/go-bindata -nomemcopy \
		-o pkg/minikube/assets/assets.go -pkg assets deploy/addons/... && \
		${LOCALBASE}/bin/go-bindata -nomemcopy \
		-o pkg/minikube/translate/translations.go -pkg translate translations/... && \
		${LOCALBASE}/bin/gofmt -s -w pkg/minikube/translate/translations.go || :

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
