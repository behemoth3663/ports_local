# Created by: Gabriel M. Dutra <0xdutra@gmail.com>

PORTNAME=	terragrunt
DISTVERSIONPREFIX=	v
DISTVERSION=	0.31.3
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

_OPTIONS_extract=	501:pre-post-extract

pre-post-extract: .PHONY
	${LN} -s ${WRKDIR}/go-commons-0.7.0 ${WRKDIR}/gruntwork-cli-0.7.0

.-include "${.CURDIR}/Makefile.deps"

#		gruntwork-io:gruntwork-cli:v0.7.0:gruntwork_io_gruntwork_cli/vendor/github.com/gruntwork-io/gruntwork-cli \
#		gruntwork-io:go-commons:v0.7.0:gruntwork_io_go_commons/vendor/github.com/gruntwork-io/gruntwork-cli \
GH_TUPLE:=	${GH_TUPLE:C/gruntwork([-_])cli/go\1commons/g}

.include <bsd.port.mk>
