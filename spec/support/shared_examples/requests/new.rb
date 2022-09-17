# frozen_string_literal: true

RSpec.shared_examples 'new.js' do
  it { is_expected.to render_template :new }

  it { expect(response.content_type).to eq 'text/javascript; charset=utf-8' }
end

RSpec.shared_examples 'new.html' do
  it { is_expected.to render_template :new }

  it { expect(response.content_type).to eq 'text/html; charset=utf-8' }
end
