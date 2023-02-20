# Created by: Danilo Egea Gondolfo <danilo@FreeBSD.org>

PORTNAME=	kops
PORTVERSION=	1.25.3
DISTVERSIONPREFIX=	v
PORTREVISION=	3
CATEGORIES=	local sysutils

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	Kubernetes Operations
WWW=		https://github.com/kubernetes/kops

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

#-kubernetes-sigs:kustomize:    v0.12.1:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api
#-kubernetes-sigs:kustomize:    v0.13.9:kubernetes_sigs_kustomize_1/vendor/sigs.k8s.io/kustomize/kyaml
#+kubernetes-sigs:kustomize:api/v0.12.1:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api
#+kubernetes-sigs:kustomize:kyaml/v0.13.9:kubernetes_sigs_kustomize_1/vendor/sigs.k8s.io/kustomize/kyaml
#                              \1                          \2    \3                                     \4
GH_TUPLE:=	${GH_TUPLE:C|^(kubernetes-sigs:kustomize:)([^:]+)(:.*/vendor/sigs\\.k8s\\.io/kustomize/)([^[:space:]]+)|\1\4/\2\3\4|}

.include <bsd.port.mk>
