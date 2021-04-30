# $FreeBSD$

PORTNAME=	terrascan
PORTVERSION=	1.5.1
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

_OPTIONS_extract=	701:post-post-extract

post-post-extract: .PHONY
	${MV} ${WRKSRC}/vendor/sigs.k8s.io/kustomize/api/api/* ${WRKSRC}/vendor/sigs.k8s.io/kustomize/api
	${MV} ${WRKSRC}/vendor/sigs.k8s.io/kustomize/kyaml/kyaml/* ${WRKSRC}/vendor/sigs.k8s.io/kustomize/kyaml

fetch-version: .PHONY get-version

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
