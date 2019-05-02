Rails.application.routes.draw do
  resources :items, only: [:create, :edit, :update, :destroy]

  constraints :year => /\d{4}/, :month => /\d{1,2}/, :day => /\d{1,2}/, :category => /[a-z_-]+/ do
    get '/(:year(/:month)(/:category))' => 'items#index'
  end

  resources :cashes, :categories

  resource :backoffice, only: :show, controller: :backoffice

  resources :consolidations, only: :index

  root to: 'items#index'
end
