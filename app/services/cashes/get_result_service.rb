# frozen_string_literal: true

module Cashes
  class GetResultService < ApplicationService
    def initialize(action_name, params)
      @action_name = action_name

      @params = params
    end

    def call
      return ::Cashes::UpdateService.call(@params) if update?

      return ::Cashes::InitializeResourceService.call if new?

      ::Cashes::GetResourceService.call(@params)
    end

    private

    def update?
      @action_name == 'update'
    end

    def new?
      @action_name == 'new'
    end
  end
end
