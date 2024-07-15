# frozen_string_literal: true

Rails.application.routes.draw do
  ActiveAdmin.routes(self)

  scope '/(:currency)' do
    resources :items, only: %i[create edit update destroy]

    resources :cashes, only: %i[index edit update]

    resources :categories

    get '(/:month)(/:category_id)', to: 'items#index', as: :root
  end
end
