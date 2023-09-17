# Simple Payment Service

- This repository contains a simple payment API that enables users to perform transactions, manage wallets, and transfer funds between accounts. The API is designed to accommodate two types of users: store users and normal consumers. Both user types have wallets but only normal consumer can make transfers between them.

## Business logic

- Only has two types off users, store users and normal consumers, both was wallets and can make transfer between than.

  ### Wallet

  - **Transfer Validation**: When a transfer is initiated, the API checks if the user has sufficient funds in their wallet to complete the transaction. Before transfer is necessary make request for other service for check if this transaction is valid.

  - **Revert Transfers**: All transfers can be reverted if any inconsistencies or issues are detected. This ensures the security and accuracy of transactions.

  - **Notification on Receipt**: When a store user receives a transfer, the API is capable of sending notifications through email or SMS to notify the recipient. However, please note that this notification service may not always be available or may experience delays

  ### Users

  - **Unique Registration**: Only one user can register on the platform with the same CPF/CNPJ (Brazilian taxpayer ID) and email. This uniqueness constraint helps maintain data integrity and prevent duplicate registrations.

  - **Consumer Transactions**: Consumers can initiate transfers among themselves and also send transfers to store users.

  - **Store User Transactions**: Store users can only receive transfers; they cannot initiate transfers to other users.

## Architecture

- The Ports and Adapters architecture, often referred to as Hexagonal Architecture, is a design pattern that emphasizes the separation of concerns and the isolation of your application's core business logic from external systems and dependencies. It achieves this separation through the use of distinct architectural components: Ports, Adapters, and the Core.

- **Core (Business Logic)**: This is the central part of your application where the core business logic resides. It represents the heart of your system and contains all the use cases, domain entities, and rules specific to your application's functionality. The core should be completely decoupled from external dependencies, such as databases, frameworks, or third-party services. It operates in isolation, making it highly testable and easier to understand.

- **Ports**: Ports are interfaces or contracts that define how the core interacts with the external world. They are essentially gateways through which the core communicates with external systems.

- **Adapters**: Adapters are concrete implementations of the ports, bridging the gap between the core and external systems
