# Created by: Danilo Egea Gondolfo <danilo@FreeBSD.org>

PORTNAME=	aws-iam-authenticator
PORTVERSION=	0.5.9
DISTVERSIONPREFIX=	v
PORTREVISION=	1
CATEGORIES=	local security

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	Use AWS IAM credentials to authenticate to a Kubernetes cluster

LICENSE=	APACHE20
WWW=		https://github.com/kubernetes-sigs/aws-iam-authenticator

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	kubernetes-sigs
GH_SUBDIR=	src/github.com/kubernetes-sigs/${PORTNAME}

GO_TARGET=	${GO_PKGNAME}/cmd/${PORTNAME}
GO_BUILDFLAGS=	-ldflags="-X main.version=${PORTVERSION}"

PLIST_FILES=	bin/${PORTNAME}

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

post-extract: .PHONY
	${LN} -s ${WRKSRC}/src/github.com/kubernetes-sigs  ${WRKSRC}/vendor/github.com/

.include <bsd.port.mk>
