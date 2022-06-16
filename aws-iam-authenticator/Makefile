# Created by: Danilo Egea Gondolfo <danilo@FreeBSD.org>

PORTNAME=	aws-iam-authenticator
PORTVERSION=	0.5.8
DISTVERSIONPREFIX=	v
CATEGORIES=	local security

MAINTAINER=	danilo@FreeBSD.org
COMMENT=	Use AWS IAM credentials to authenticate to a Kubernetes cluster

LICENSE=	APACHE20

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	kubernetes-sigs
GH_SUBDIR=	src/github.com/kubernetes-sigs/${PORTNAME}

GO_TARGET=	${GO_PKGNAME}/cmd/${PORTNAME}
GO_BUILDFLAGS=	-ldflags="-X main.version=${PORTVERSION}"

PLIST_FILES=	bin/${PORTNAME}

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
