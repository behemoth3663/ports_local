# $FreeBSD$

PORTNAME=	terrascan
PORTVERSION=	1.4.0
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

PLIST_FILES=	bin/${PORTNAME}

fetch-version: .PHONY get-version

_OPTIONS_extract=	701:post-post-extract

post-post-extract: .PHONY
	${MV} ${WRKSRC}/vendor/sigs.k8s.io/kustomize/api/api/* ${WRKSRC}/vendor/sigs.k8s.io/kustomize/api
	${MV} ${WRKSRC}/vendor/sigs.k8s.io/kustomize/kyaml/kyaml/* ${WRKSRC}/vendor/sigs.k8s.io/kustomize/kyaml

#rm -f -- distinfo Makefile.deps && make makesum && make gomod-vendor && make gomod-vendor >Makefile.deps && make makesum && mkcal && instpc
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
