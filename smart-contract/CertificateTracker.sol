pragma solidity ^0.6.4;

contract CertificateTracker {

    struct CertificateOwner  {
        address certificateOwnerAddress;
        string name;
        string email;
    }

    event CertificateIssued (
        address issuerAddress,
        address ownerAddress,
        string name,
        string email,
        string certificate
    );

    mapping (string => CertificateOwner) Certificate;
    mapping (address => mapping (address => string)) Issuer;

    function setCertificateDetails(address _ownerAddress, string memory _name, string memory _email, string memory _certificate) public {
        Certificate[_certificate].certificateOwnerAddress = _ownerAddress;
        Certificate[_certificate].name = _name;
        Certificate[_certificate].email = _email;
        Issuer[msg.sender][_ownerAddress] = _certificate;
        emit CertificateIssued(msg.sender, _ownerAddress, _name, _email, _certificate);
    }

    function getOwnerDetailsByCertificateHash(string memory _certificate) public view returns (address, string memory, string memory) {
        return (Certificate[_certificate].certificateOwnerAddress, Certificate[_certificate].name, Certificate[_certificate].email);
    }

    function getIssuedCertificateDetails(address _ownerAddress) public view returns (address, string memory, string memory) {
        return (Certificate[Issuer[msg.sender][_ownerAddress]].certificateOwnerAddress, Certificate[Issuer[msg.sender][_ownerAddress]].name, Certificate[Issuer[msg.sender][_ownerAddress]].email);
    }

    function getCertificateHash(address _issuerAddress, address _ownerAddress) public view returns (string memory) {
        return Issuer[_issuerAddress][_ownerAddress];
    }

}
