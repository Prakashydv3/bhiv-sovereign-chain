// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract AnchorContract {
    struct Anchor {
        bytes32 anchorId;
        string artifactType;
        bytes32 artifactHash;
        bytes32 parentHash;
        uint256 timestampUtc;
        uint8 schemaVersion;
    }

    mapping(bytes32 => Anchor) private anchors;
    
    event AnchorCreated(
        bytes32 indexed anchorId,
        string artifactType,
        bytes32 artifactHash,
        bytes32 parentHash,
        uint256 timestampUtc,
        uint8 schemaVersion
    );

    function createAnchor(
        string memory _artifactType,
        bytes32 _artifactHash,
        bytes32 _parentHash,
        uint8 _schemaVersion
    ) external returns (bytes32) {
        require(bytes(_artifactType).length > 0, "artifact_type required");
        require(_artifactHash != bytes32(0), "artifact_hash required");
        require(_schemaVersion > 0, "schema_version required");

        uint256 timestamp = block.timestamp;
        bytes32 anchorId = keccak256(abi.encodePacked(
            _artifactType,
            _artifactHash,
            _parentHash,
            timestamp,
            _schemaVersion,
            msg.sender
        ));

        anchors[anchorId] = Anchor({
            anchorId: anchorId,
            artifactType: _artifactType,
            artifactHash: _artifactHash,
            parentHash: _parentHash,
            timestampUtc: timestamp,
            schemaVersion: _schemaVersion
        });

        emit AnchorCreated(
            anchorId,
            _artifactType,
            _artifactHash,
            _parentHash,
            timestamp,
            _schemaVersion
        );

        return anchorId;
    }

    function getAnchor(bytes32 _anchorId) external view returns (
        bytes32 anchorId,
        string memory artifactType,
        bytes32 artifactHash,
        bytes32 parentHash,
        uint256 timestampUtc,
        uint8 schemaVersion
    ) {
        Anchor memory anchor = anchors[_anchorId];
        require(anchor.anchorId != bytes32(0), "Anchor does not exist");
        
        return (
            anchor.anchorId,
            anchor.artifactType,
            anchor.artifactHash,
            anchor.parentHash,
            anchor.timestampUtc,
            anchor.schemaVersion
        );
    }
}
