#

PORTNAME=	awsets
PORTVERSION=	1.0.4
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

GH_COMMIT=	git ls-remote --tags \
		https://github.com/${GH_ACCOUNT}/${PORTNAME} \
		refs/tags/${DISTVERSIONPREFIX}${PORTVERSION}

GO_BUILDFLAGS+=	-ldflags="-s -w \
	-X 'main.version=${DISTVERSIONPREFIX}${PORTVERSION}' \
	-X 'main.commit=${:!${GH_COMMIT}!:[1]:C/(.{7}).*/\1/}' \
	-X 'main.date=${%FT%TZ:L:gmtime}'"

fetch-version: .PHONY get-version

post-patch: .PHONY
	${MKDIR} ${WRKSRC}/vendor/github.com/${GH_ACCOUNT}/${PORTNAME}
	${FIND} ${WRKSRC}/../.. -maxdepth 1 \( \
		\( -type d -not -name '.*' -not -name cmd \) -or \
		\( -type f \( -name '*.go' -or -name 'go.*' \) \) \
	\) -exec ${RLN} {} ${WRKSRC}/vendor/github.com/${GH_ACCOUNT}/${PORTNAME} \;

#	cd ${WRKSRC}/vendor/github.com/${GH_ACCOUNT}/${PORTNAME}/lister && \
#	${REINPLACE_CMD} -r 's|MaxResults:[[:space:]]+aws\.|MaxResults: *aws.|' \
#		ec2_dhcpoption.go ec2_flowlog.go ec2_instance.go ec2_internetgateway.go \
#		ec2_launchtemplate.go ec2_natgateway.go ec2_networkacl.go \
#		ec2_networkinterface.go ec2_routetable.go ec2_securitygroup.go \
#		ec2_snapshot.go ec2_spotfleet.go ec2_subnet.go ec2_transitgateway.go \
#		ec2_volume.go ec2_vpc.go ec2_vpcendpoint.go ec2_vpcendpointservice.go \
#		ec2_vpcpeering.go && \
#	${REINPLACE_CMD} -r 's|AllRegions:[[:space:]]+aws\.|AllRegions: *aws.|' ../main.go

.-include "${.CURDIR}/Makefile.deps"

.include <bsd.port.mk>
