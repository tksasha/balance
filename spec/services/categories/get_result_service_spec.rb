# frozen_string_literal: true

RSpec.describe Categories::GetResultService do
  subject { described_class.new action_name, params }

  let(:params) { double }

  describe '#call' do
    context do
      let(:action_name) { 'new' }

      before { allow(Categories::InitializeService).to receive(:call).and_return(:result) }

      its(:call) { should eq :result }
    end

    context do
      let(:action_name) { 'create' }

      before { allow(Categories::CreateService).to receive(:call).and_return(:result) }

      its(:call) { should eq :result }
    end
  end
end
