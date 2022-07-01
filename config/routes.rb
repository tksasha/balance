# frozen_string_literal: true

Rails.application.routes.draw do
  scope '/(:currency)' do
    resources :items, only: %i[create edit update destroy]

    resources :cashes, only: %i[index edit update]

    resource :backoffice, only: :show, controller: 'backoffice/dashboard'

    resources :consolidations, only: :index

    namespace :backoffice do
      resources :cashes, except: :show

      resources :categories
    end

    get '(/:month)(/:category_id)', to: 'items#index', as: :root
  end
end
