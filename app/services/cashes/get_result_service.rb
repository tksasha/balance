# frozen_string_literal: true

module Cashes
  class GetResultService < ApplicationService
    def initialize(action_name, params)
      @action_name = action_name

      @params = params
    end

    def call
      return ::Cashes::InitializeService.call(@params) if new?

      return ::Cashes::GetResourceService.call(@params) if show? || edit?

      "::Cashes::#{ @action_name.camelize }Service".constantize.call(@params)
    end

    private

    def new?
      @action_name == 'new'
    end

    def show?
      @action_name == 'show'
    end

    def edit?
      @action_name == 'edit'
    end
  end
end
