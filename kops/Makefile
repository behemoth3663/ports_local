# Created by: Danilo Egea Gondolfo <danilo@FreeBSD.org>

PORTNAME=	kops
PORTVERSION=	1.29.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	5
CATEGORIES=	local sysutils

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	Kubernetes Operations
WWW=		https://github.com/${GH_ACCOUNT}/${GH_PROJECT}

LICENSE=	APACHE20

BUILD_DEPENDS=	go-bindata:devel/go-bindata

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	kubernetes

GO_PKGNAME=	k8s.io/${PORTNAME}
GO_TARGET=	${GO_PKGNAME}/cmd/${PORTNAME}

GO_BUILDFLAGS+=	-ldflags="-s -w -extldflags '-static' \
	-X k8s.io/kops.Version=${PORTVERSION}"

PLIST_FILES=	bin/${PORTNAME}

pre-build: .PHONY
	cd ${WRKSRC} && \
		${LOCALBASE}/bin/go-bindata \
			-o upup/models/bindata.go -pkg models \
			-prefix upup/models/ upup/models/...

pre-deps pre-deps-all post-deps-all: .PHONY

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.sinclude "${.CURDIR}/Makefile.deps"

#-kubernetes-sigs:kustomize:    v0.12.1:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api
#-kubernetes-sigs:kustomize:    v0.13.9:kubernetes_sigs_kustomize_1/vendor/sigs.k8s.io/kustomize/kyaml
#+kubernetes-sigs:kustomize:api/v0.12.1:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api
#+kubernetes-sigs:kustomize:kyaml/v0.13.9:kubernetes_sigs_kustomize_1/vendor/sigs.k8s.io/kustomize/kyaml
GH_TUPLE:=	${GH_TUPLE:C|^(kubernetes-sigs:kustomize:)([^:]+)(:.*/vendor/sigs\\.k8s\\.io/kustomize/)([^[:space:]]+)|\1\4/\2\3\4|}
#		kubernetes-sigs:kustomize:6ce0bf390ce3:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api \
#		kubernetes-sigs:kustomize:6ce0bf390ce3:kubernetes_sigs_kustomize_1/vendor/sigs.k8s.io/kustomize/kyaml \

#-googleapis:google-cloud-go:                 v0.2.3:googleapis_google_cloud_go/vendor/cloud.google.com/go/compute/metadata
#+googleapis:google-cloud-go:compute/metadata/v0.2.3:googleapis_google_cloud_go/vendor/cloud.google.com/go/compute/metadata
GH_TUPLE:=	${GH_TUPLE:C,^(googleapis:google-cloud-go):(.*)/(compute/metadata|auth/oauth2adapt)$,\1:\3/\2,}

.include <bsd.port.mk>
