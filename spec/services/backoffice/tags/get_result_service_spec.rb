# frozen_string_literal: true

RSpec.describe Backoffice::Tags::GetResultService do
  subject { described_class.new action_name, category, params }

  let(:category) { build :category }

  let(:params) { double }

  describe '#call' do
    context do
      let(:action_name) { 'new' }

      before { allow(Backoffice::Tags::InitializeService).to receive(:call).with(category).and_return(:result) }

      its(:call) { should eq :result }
    end

    context do
      let(:action_name) { 'create' }

      before { allow(Backoffice::Tags::CreateService).to receive(:call).with(category, params).and_return(:result) }

      its(:call) { should eq :result }
    end

    context do
      let(:action_name) { 'edit' }

      before { allow(Backoffice::Tags::GetResourceService).to receive(:call).with(params).and_return(:result) }

      its(:call) { should eq :result }
    end

    context do
      let(:action_name) { 'update' }

      before { allow(Backoffice::Tags::UpdateService).to receive(:call).with(params).and_return(:result) }

      its(:call) { should eq :result }
    end
  end
end
