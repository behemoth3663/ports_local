# Created by: Gabriel M. Dutra <0xdutra@gmail.com>

PORTNAME=	terragrunt
DISTVERSION=	0.30.3
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	0xdutra@gmail.com
COMMENT=	Terragrunt is a thin wrapper for Terraform

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

USES=	go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	gruntwork-io

GO_BUILDFLAGS=	-ldflags="-X main.VERSION=v${PORTVERSION}"

PLIST_FILES=	bin/${PORTNAME}

.-include "${.CURDIR}/Makefile.deps"

#		gruntwork-io:gruntwork-cli:v0.7.0:gruntwork_io_gruntwork_cli/vendor/github.com/gruntwork-io/gruntwork-cli \
#		gruntwork-io:go-commons:v0.7.0:gruntwork_io_go_commons/vendor/github.com/gruntwork-io/gruntwork-cli \
GH_TUPLE:=	${GH_TUPLE:C/gruntwork([-_])cli/go\1commons/g}

.include <bsd.port.mk>
