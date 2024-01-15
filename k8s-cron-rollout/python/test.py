import argparse

def parse_args():
    parser = argparse.ArgumentParser(description='Rollout a k8s deployment')
    parser.add_argument('-a', '--auth', type=str, required=True, choices=["inside", "outside"],
        help="Specify auth method based on RBAC (inside a pod) or kubeconfig (outside the cluster).")
    parser.add_argument('-ns', '--namespace', type=str, help="K8s namespace.")
    parser.add_argument('-deploy', '--deployment-name', type=str, help="Deployment name to rollout.")
    parser.add_argument('-cc', '--change-cause', type=str, help="Change cause for annotation kubernetes.io/change-cause")
    args = parser.parse_args()
    return args

def main():
    args = parse_args()
    if args.auth == "inside":
        print(args.auth)
    elif args.auth == "outside":
        print(args.auth)
    test = args.change_cause if args.change_cause else "cronjob execution"
    print(test)


if __name__ == "__main__":
    main()
