#

PORTNAME=	terrascan
PORTVERSION=	1.8.1
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Compliance and security violations in terraform code

LICENSE=	APACHE20
LICENSE_FILE=	${WRKSRC}/LICENSE

USES=		go:modules
GO_TARGET=	./cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	accurics
GH_TAGNAME=	d8fd9c4

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

_OPTIONS_extract=	701:post-post-extract

post-post-extract: .PHONY
	${MV} ${WRKSRC}/vendor/sigs.k8s.io/kustomize/api/api/* ${WRKSRC}/vendor/sigs.k8s.io/kustomize/api
	${MV} ${WRKSRC}/vendor/sigs.k8s.io/kustomize/kyaml/kyaml/* ${WRKSRC}/vendor/sigs.k8s.io/kustomize/kyaml

.-include "${.CURDIR}/Makefile.deps"

#-               kubernetes-sigs:kustomize:v0.10.15:kubernetes_sigs_kustomize_1/vendor/sigs.k8s.io/kustomize/kyaml \
#-               kubernetes-sigs:kustomize:v0.8.5:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api \
#+               kubernetes-sigs:kustomize:kyaml/v0.10.15:kubernetes_sigs_kustomize_1/vendor/sigs.k8s.io/kustomize/kyaml \
#+               kubernetes-sigs:kustomize:api/v0.8.5:kubernetes_sigs_kustomize/vendor/sigs.k8s.io/kustomize/api \

.include <bsd.port.mk>
