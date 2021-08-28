#

PORTNAME=	awsets
PORTVERSION=	1.0.2
DISTVERSIONPREFIX=	v
CATEGORIES=	sysutils

MAINTAINER=	none@freebsd.org
COMMENT=	Utility for export all AWS resources

LICENSE=	MIT
LICENSE_FILE=	${WRKSRC}/../../LICENSE

RUN_DEPENDS=	git:devel/git

USES=		go:modules

WRKSRC_SUBDIR=	cmd/${PORTNAME}

USE_GITHUB=	yes
GH_ACCOUNT=	trek10inc

PLIST_FILES=	bin/${PORTNAME}

#8d9d922d9dc592bedd5ac77b89c71a5f5b1b7e1c	refs/tags/v0.8.0
GH_SHORT_HASH=	git ls-remote --tags https://github.com/${GH_ACCOUNT}/${PORTNAME}.git \
	| ${SED} -rn 's|^(.{7}).*/${DISTVERSIONPREFIX}${PORTVERSION:S/./\\&/g}$$|\1|p'

GO_BUILDFLAGS+=	-ldflags="-s \
	-X 'main.version=${PORTVERSION}' \
	-X 'main.commit=${:!${GH_SHORT_HASH}!}' \
	-X 'main.date=${:!date +%F_%T!:S/_/ /}'"

fetch-version: .PHONY get-version

post-patch: .PHONY
	${MKDIR} ${WRKSRC}/vendor/github.com/${GH_ACCOUNT}/${PORTNAME}
	${FIND} ${WRKSRC}/../.. -maxdepth 1 \( \
		\( -type f -and -name '*.go' \) -or \
		\( -type d -and -not -name '.*' -and -not -name 'cmd' \) \
	\) -exec ${RLN} {} ${WRKSRC}/vendor/github.com/${GH_ACCOUNT}/${PORTNAME} \;
	cd ${WRKSRC}/vendor/github.com/${GH_ACCOUNT}/${PORTNAME}/lister && \
	${REINPLACE_CMD} -r 's|MaxResults:[[:space:]]+aws\.|MaxResults: *aws.|' \
		ec2_dhcpoption.go ec2_flowlog.go ec2_instance.go ec2_internetgateway.go \
		ec2_launchtemplate.go ec2_natgateway.go ec2_networkacl.go \
		ec2_networkinterface.go ec2_routetable.go ec2_securitygroup.go \
		ec2_snapshot.go ec2_spotfleet.go ec2_subnet.go ec2_transitgateway.go \
		ec2_volume.go ec2_vpc.go ec2_vpcendpoint.go ec2_vpcendpointservice.go \
		ec2_vpcpeering.go && \
	${REINPLACE_CMD} -r 's|AllRegions:[[:space:]]+aws\.|AllRegions: *aws.|' ../main.go

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
