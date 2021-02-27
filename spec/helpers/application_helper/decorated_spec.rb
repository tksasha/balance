# frozen_string_literal: true

RSpec.describe ApplicationHelper, type: :helper do
  subject { helper }

  describe '#decorated' do
    before { allow(helper).to receive_message_chain(:resource, :decorate).and_return(:decorated) }

    subject { helper.decorated }

    it { should eq :decorated }
  end
end
