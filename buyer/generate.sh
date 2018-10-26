#!/bin/bash
swagger generate client -f ../third_party/equinix/ecx/ecxfabric-v3-catalog-buyer.json -A go-ecxfabric/buyer
# apply o_auth tokentimeout patch
patch models/o_auth_response.go < buyer_patch.diff
