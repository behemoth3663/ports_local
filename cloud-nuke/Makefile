#

PORTNAME=	cloud-nuke
PORTVERSION=	0.31.0
DISTVERSIONPREFIX=	v
#PORTREVISION=	1
CATEGORIES=	local sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	CLI tool to delete all resources in an AWS account
WWW=		https://github.com/${GH_ACCOUNT}/${PORTNAME}

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/LICENSE.txt

USES=		go:modules

USE_GITHUB=	yes
GH_ACCOUNT=	gruntwork-io

PLIST_FILES=	bin/${PORTNAME}

_OPTIONS_extract=	501:pre-post-extract
#Resolve the error:
#mv: rename /usr/obj/usr/ports/local/cloud-nuke/work/gruntwork-cli-0.7.0 to /usr/obj/usr/ports/local/cloud-nuke/work/cloud-nuke-0.29.4/vendor/github.com/gruntwork-io/gruntwork-cli: No such file or directory
#/usr/obj/usr/ports/local/cloud-nuke/work/gruntwork-cli-0.7.0 -> /usr/obj/usr/ports/local/cloud-nuke/work/go-commons-0.7.0
#/usr/obj/usr/ports/local/cloud-nuke/work/gruntwork-cli-0.16.1 -> /usr/obj/usr/ports/local/cloud-nuke/work/go-commons-0.16.1
pre-post-extract: .PHONY .SILENT
	test -z '${WRKDIR}' || ${FIND} ${WRKDIR} -maxdepth 1 -type d -name 'go-commons-*' | while read src; do \
		${LN} -s "$${src}" "$${src%%go-commons-*}gruntwork-cli-$${src##*go-commons-}"; \
	done

.include "${COMPAT_PREFIX}/etc/make.d/go-deps.mk"
.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
