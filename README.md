# Go_password_1
Built over 3 days as a university project to explore cryptography and secure programming in Go.

# Password Manager CLI

A command-line password manager built in Go with AES-256 encryption for secure password storage and management.

## ⚠️ Educational Project Notice

This project was built as a learning exercise to understand cryptography, encryption, and secure password management in Go. **This is not intended for production use.** For managing real passwords, please use established solutions like 1Password, Bitwarden, or similar professionally audited tools.

## Features

- **AES-256-CBC Encryption**: All passwords and site names are encrypted before storage
- **Secure Random Generation**: Cryptographically secure password generation using `crypto/rand`
- **PKCS7 Padding**: Proper block cipher padding implementation
- **CRUD Operations**: Create, read, update, and delete password entries
- **Search Functionality**: Search through encrypted passwords by site name
- **Color-Coded CLI**: User-friendly terminal interface with color formatting
- **JSON Storage**: Encrypted data persisted to JSON file

## Project Structure

```
.
├── main.go                 # Application entry point and menu logic
├── counter/
│   └── counter.go         # First-run detection mechanism
├── display/
│   └── display.go         # Terminal UI and formatting
├── encrypt/
│   └── enndecrypt.go      # AES encryption/decryption with PKCS7 padding
├── key/
│   └── keygenerate.go     # Cryptographically secure key generation
├── password/
│   └── password.go        # Password generation logic
├── storage/
│   └── filestorage.go     # JSON file persistence
└── utils/
    └── constants.go       # ANSI color codes for terminal output
```

## Technical Implementation

### Encryption
- **Algorithm**: AES-256 in CBC mode
- **Key Size**: 256 bits (32 bytes)
- **IV**: Random 16-byte initialization vector generated per encryption
- **Padding**: PKCS7 padding for block alignment
- **Storage Format**: Hex-encoded ciphertext with prepended IV

### Password Generation
- **Character Space**: 95 printable ASCII characters (letters, numbers, symbols)
- **Default Length**: 14 characters
- **Randomness**: Uses `crypto/rand` with `big.Int` for unbiased selection

## Installation & Usage

```bash
# Clone the repository
git clone https://github.com/yourusername/password-manager-go
cd password-manager-go

# Run the application
go run main.go
```

### First Run
On first launch, the application generates a 32-character encryption key. **Save this key securely** - you'll need it to decrypt your passwords in future sessions.

### Menu Options
1. View all passwords
2. Add new password
3. Add new password and view all
4. Update existing password
5. Delete existing password
6. Search existing password
7. Exit

## What I Learned

This project taught me:

- **Cryptography Fundamentals**: Implementing AES encryption, understanding block ciphers, IV generation, and padding schemes
- **Security Best Practices**: Using `crypto/rand` instead of `math/rand`, proper key management considerations
- **Go Package Design**: Organizing code into logical packages with clear responsibilities
- **Error Handling**: Comprehensive error handling and propagation patterns in Go
- **File I/O**: Reading and writing JSON data with proper error checking
- **CLI Design**: Creating user-friendly terminal interfaces with color coding

## Known Limitations & Future Improvements

### Current Limitations
- **Key Management**: Requires users to remember/store a 32-character random string (impractical)
- **First-Run Detection**: Uses a counter file which is fragile and not intuitive
- **No Key Derivation**: Keys are generated randomly rather than derived from a user password
- **No Tests**: Project lacks unit tests for critical cryptographic functions
- **Single User**: No multi-user support or access control
- **No Backup**: No built-in mechanism for secure backups

### Planned Improvements (v2)
- [ ] **Argon2 Key Derivation**: Allow users to enter a memorable password, derive encryption key using Argon2id
- [ ] **Salt Storage**: Store salt in configuration file for consistent key derivation
- [ ] **Replace Counter Mechanism**: Use existence of salt/config file to detect first run
- [ ] **Unit Tests**: Add comprehensive tests for encryption, decryption, and padding
- [ ] **Password Strength Validation**: Check and warn about weak generated passwords
- [ ] **Export/Import**: Allow secure backup and restoration of password database
- [ ] **Master Password Change**: Ability to re-encrypt all passwords with a new master password
- [ ] **Clipboard Integration**: Copy passwords to clipboard with auto-clear

## Security Considerations

**Why this project should not be used in production:**

1. **No Security Audit**: This code has not been professionally audited
2. **Key Management**: Current key storage approach is insecure
3. **No Rate Limiting**: Vulnerable to brute-force attacks on the master key
4. **Timing Attacks**: Decryption timing could leak information
5. **Memory Safety**: Sensitive data not securely wiped from memory
6. **No Integrity Checks**: Missing HMAC or authenticated encryption (AES-GCM would be better)

## Development Timeline

Built over 3 days as a university project to explore cryptography and secure programming in Go.

## Reflections

Building this project taught me that implementing cryptography correctly is **hard**. While the code works functionally, production password managers require:
- Professional security audits
- Threat modeling
- Secure memory handling
- Defense against side-channel attacks
- Robust key derivation
- Regular security updates

This experience deepened my respect for projects like 1Password and Bitwarden, and reinforced why cryptography should be left to experts whenever possible.

## License

MIT License - Feel free to use this for learning purposes.

## Acknowledgments

Built as a learning project to understand:
- Go's `crypto` standard library
- AES encryption implementation
- Secure password management principles
- CLI application design patterns

---

**Built by**: Micheal Dand
**Year**: 2024  
**Context**: University project exploring cryptography and secure systems
