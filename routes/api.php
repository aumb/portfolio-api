<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\GeneralController;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::post('education', [GeneralController::class, 'postEducation']);
Route::post('personal_information', [GeneralController::class, 'postPersonalInformation']);
Route::post('job', [GeneralController::class, 'postJob']);
Route::post('recentwork', [GeneralController::class, 'postRecentWork']);

Route::get('information', [GeneralController::class, 'getInformation']);
Route::get('profilepicture/{id}', [GeneralController::class, 'getProfilePicture']);
Route::get('recentworkpicture/{id}', [GeneralController::class, 'getRecentWorkPicture']);
Route::get('jobpicture/{id}', [GeneralController::class, 'getJobPicture']);

Route::delete('personal_information', [GeneralController::class, 'deletePersonalInformation']);
Route::delete('job', [GeneralController::class, 'deleteJob']);
Route::delete('recentwork', [GeneralController::class, 'deleteRecentWork']);

Route::middleware('auth:api')->get('/user', function (Request $request) {
    return $request->user();
});

Route::fallback(function () {
    return response()->json([
        'message' => 'Page Not Found. If error persists, contact info@website.com'
    ], 404);
});
