# frozen_string_literal: true

module Cashes
  class GetResultService < ApplicationService
    def initialize(action_name, params)
      @action_name = action_name

      @params = params
    end

    def call
      return ::Cashes::UpdateService.call(@params) if update?

      return ::Cashes::InitializeService.call if new?

      return ::Cashes::CreateService.call(@params) if create?

      ::Cashes::GetResourceService.call(@params)
    end

    private

    def update?
      @action_name == 'update'
    end

    def new?
      @action_name == 'new'
    end

    def create?
      @action_name == 'create'
    end
  end
end
