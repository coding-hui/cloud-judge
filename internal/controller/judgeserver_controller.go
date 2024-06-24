/*
Copyright 2024 coding-hui.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	judgev1 "github.com/coding-hui/cloud-judge/api/v1"
)

// JudgeServerReconciler reconciles a JudgeServer object
type JudgeServerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=judge.wecoding.top,resources=judgeservers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=judge.wecoding.top,resources=judgeservers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=judge.wecoding.top,resources=judgeservers/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// Modify the Reconcile function to compare the state specified by
// the JudgeServer object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.2/pkg/reconcile
func (r *JudgeServerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := ctrl.LoggerFrom(ctx)
	var err error

	judgeServer := &judgev1.JudgeServer{}
	if err = r.Get(ctx, req.NamespacedName, judgeServer); err == nil {
		return r.judgeServerReconcile(ctx, req, judgeServer)
	}

	// No match found
	if apierrors.IsNotFound(err) {
		logger.Info("Read request judge server not found error!")
	} else {
		logger.Error(err, "Read request judge server error!")
	}

	// Error reading the object - requeue the request.
	return ctrl.Result{}, client.IgnoreNotFound(err)
}

// SetupWithManager sets up the controller with the Manager.
func (r *JudgeServerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&judgev1.JudgeServer{}).
		Complete(r)
}

func (r *JudgeServerReconciler) judgeServerReconcile(ctx context.Context, req ctrl.Request, judgeServer *judgev1.JudgeServer) (ctrl.Result, error) {
	return ctrl.Result{}, nil
}
