<?php

namespace App\Http\Controllers;

use Exception;
use App\Http\Resources\InformationResource;
use App\Models\Information;

class GeneralController extends Controller
{
    public function getInformation()
    {

        try {
            $information = new Information;
            return new InformationResource($information);
        } catch (Exception $e) {
            return response()->json(['error' => $e->getMessage()], 400);
        }
    }
}
